const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const express = require('express');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

const PROTO_PATH = __dirname + '/messaging.proto';

const clients = [];
const storedMessages = [];

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true,
    }
);
const messagingProto = grpc.loadPackageDefinition(packageDefinition).messaging;

const client = new messagingProto.Messaging('go:8080', grpc.credentials.createInsecure());

app.use((req, res, next) => {
    res.setHeader('Access-Control-Allow-Origin', 'http://localhost:3000');
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
    res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
    res.setHeader('Access-Control-Allow-Credentials', 'true');
    next();
});


// Custom middleware to add res.flush() method
app.use((req, res, next) => {
    res.flush = () => {
        if (res.socket && res.socket.writable) {
            res.write('\n');
        }
    };
    next();
});


// Endpoint to send a message
app.post('/send-message', (req, res) => {
    const { message } = req.body;

    if (!message) {
        return res.status(400).json({ error: 'Missing message in request body' });
    }

    storedMessages.push(message); // Add the new message to the storedMessages array

    // Broadcast the new message to all connected clients
    clients.forEach(client => {
        client.write(`event: message\n`);
        client.write(`data: ${message}\n\n`);
        client.flush();
    });

    console.log('Message sent:', message);
    res.json({ message: 'Message sent successfully' });
});

// Endpoint to subscribe and print all new messages
// Endpoint to subscribe and print all new messages
app.get('/subscribe', (req, res) => {
    clients.push(res);

    res.writeHead(200, {
        'Content-Type': 'text/event-stream',
        'Cache-Control': 'no-cache',
        Connection: 'keep-alive',
        'Transfer-Encoding': 'chunked',
    });

    // Send stored messages to the client
    storedMessages.forEach((message) => {
        res.write(`event: message\n`);
        res.write(`data: ${message}\n\n`);
    });
    res.flush();

    res.write(': connected\n\n');
    res.write('retry: 5000\n\n');

    // Remove client from the list when the connection is closed
    req.on('close', () => {
        clients.splice(clients.indexOf(res), 1);
    });
});



app.listen(3535, () => {
    console.log('Server listening on port 3535');
});

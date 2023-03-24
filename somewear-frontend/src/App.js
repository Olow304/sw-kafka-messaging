import 'eventsource-polyfill';
import './App.css';
import {useEffect, useRef, useState} from "react";
import Message from "./components/Message";
import MessageView from "./components/MessageView";


let intervalId = null;

function App() {
    const [messages, setMessages] = useState([]);
    const [connectionStatus, setConnectionStatus] = useState('connecting');
    const [isNewMessage, setIsNewMessage] = useState(false);


    const handleVisibilityChange = () => {
        if (!document.hidden) {
            document.title = 'SW - DEMO';

            // Remove the click event listener to stop playing the sound when the user clicks anywhere on the page
            document.removeEventListener('click', handleClick);
        }
    };

    const handleClick = () => {
        // Play sound
        notificationSound.play();

        // Remove event listener
        document.removeEventListener('click', handleClick);
    };

    const handleMessage = (event) => {
        const newMessage = event.data;
        console.log('Received message:', newMessage);
        setMessages((prevMessages) => [...prevMessages, newMessage]);

        // Play sound and create a notification if the tab is not active and there's a new message
        if (document.hidden) {
            setIsNewMessage(true);
            if (Notification.permission === 'granted') {
                const notification = new Notification('New Message', {
                    body: newMessage,
                    icon: '/favicon.ico',
                });
            } else if (Notification.permission !== 'denied') {
                Notification.requestPermission().then((permission) => {
                    if (permission === 'granted') {
                        const notification = new Notification('New Message', {
                            body: newMessage,
                            icon: '/favicon.ico',
                        });
                    }
                });
            }


            // Add an event listener to reset the title when the tab becomes active again
            document.addEventListener('visibilitychange', handleVisibilityChange, { once: true });

            // Play sound after a short delay
            setTimeout(() => {
                notificationSound.play();
            }, 100);

            // Add event listener to document
            document.addEventListener('click', handleClick);
        }
    };

    const subscribeToMessages = () => {
        const noCacheUrl = `http://localhost:3535/subscribe?nocache=${Math.random()}`;
        const eventSource = new EventSource(noCacheUrl);

        let retryCount = 0;
        const maxRetries = 10;
        const retryDelay = 1000;

        eventSource.addEventListener('message', handleMessage);

        eventSource.onerror = (error) => {
            console.error('EventSource error:', error);
            eventSource.close();
            retryCount++;

            if (retryCount <= maxRetries) {
                setTimeout(() => {
                    console.log(`Retrying connection (attempt ${retryCount})...`);
                    subscribeToMessages();
                }, retryDelay);
            } else {
                console.error(`Max retries (${maxRetries}) exceeded, giving up`);
            }
        };

        return () => {
            eventSource.close();
        };
    };

    const notificationSound = new Audio('/notification.mp3');

    useEffect(() => {
        const unsubscribe = subscribeToMessages();
        return () => {
            unsubscribe();
        };
    }, []);

    return (
        <MessageView isNewMessage={isNewMessage} onMessageViewClick={() => setIsNewMessage(false)}>
            <Message messages={messages}/>
        </MessageView>
    );
}

export default App;


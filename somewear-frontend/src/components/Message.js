const Message = ({messages}) => {
    const dateFormat = new Date().toLocaleString();

    return (
        <div style={{paddingTop: '20px'}}>
            <span className="chat-label">hiking-legends (32 members)</span>
            <ul className="message-content">
                {messages.map((message, index) => (
                    <li className="message-item" key={index}>
                        <p className="author-date-info"><strong style={{color: 'blue'}}>demo-olow</strong> - today</p>
                        <p className="message-text">{message}</p>
                    </li>
                ))}

            </ul>
        </div>
    );
};

export default Message;
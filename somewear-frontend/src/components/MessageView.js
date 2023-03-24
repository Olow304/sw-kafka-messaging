import Navbar from "./Navbar";
import LeftNavbar from "./LeftNavbar";
import ToggleViewMembers from "./ToggleViewMembers";

import "../styles/MessageStyle.css"

let intervalId = null;

const MessageView = ({ children, isNewMessage, onMessageViewClick }) => {

    const title = isNewMessage ? "New Message" : "Original Title";

    // intervalId = setInterval(() => {
    //     document.title = isNewMessage ? "New Message" : "Original Title";
    // }, 1000);

    return (
        <div className="message-view-container" onClick={onMessageViewClick}>
            <Navbar activeTab="messages" />
            <div className="main-container">
                <LeftNavbar />
                <div className="message-members-section">
                    { children }

                </div>
            </div>
        </div>
    );
};

export default MessageView
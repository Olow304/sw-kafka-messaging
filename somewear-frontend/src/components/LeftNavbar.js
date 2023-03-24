const LeftNavbar = () => {
    return (
        <div className="leftnavbar">
            <span className="message-title">Messages</span>
            <div className="search-container">
                <input type="text" className="search-input" placeholder="Search" />
            </div>
            <div className="message-group-list">
                <div className="message-group-item active">
                    <div className="message-group-item__name">Hiking Legends</div>
                    <div className="message-group-item__date">today</div>
                </div>
                <div className="message-group-item">Group 2</div>
            </div>
        </div>
    )
}

export default LeftNavbar

const Navbar = ({ activeTab }) => {
    return (
        <div className="navbar">
            <span className="logo">SW-DEMO</span>
            <div className="navbar-right-content">
                <div className={`navbar-right-item${activeTab === 'map' ? ' active' : ''}`}>
                    <span>MAP</span>
                </div>
                <div className={`navbar-right-item${activeTab === 'messages' ? ' active' : ''}`}>
                    <span>MESSAGES</span>
                </div>
                <div className={`navbar-right-item${activeTab === 'account' ? ' active' : ''}`}>
                    <span>ACCOUNT</span>
                </div>
                <div className={`navbar-right-item${activeTab === 'profile' ? ' active' : ''}`}>
                    <span>DEMO-OLOW</span>
                </div>
            </div>
        </div>
    )
}

export default Navbar
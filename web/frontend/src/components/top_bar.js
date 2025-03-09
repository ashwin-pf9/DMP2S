import React from "react";
import { useNavigate } from "react-router-dom";
import "../styles/top_bar.css"; // Make sure this file includes the above CSS

const TopBar = () => {
  const navigate = useNavigate();

  const handleEditProfile = () => {
    navigate("/edit-profile");
  };

  const handleLogout = () => {
    // Add logout logic here (e.g., clearing tokens, redirecting)
    navigate("/login");
  };

  return (
    <div className="top-bar">
      <div className="project-name">Pipeline Management System</div>

      <div className="user-profile">
        <img
          src="/path-to-profile-pic.jpg"
          alt="Profile"
          className="profile-icon"
        />
        
        <div className="profile-dropdown">
          <a onClick={handleEditProfile}>Edit Profile</a>
          <a onClick={handleLogout}>Log Out</a>
        </div>
      </div>
    </div>
  );
};

export default TopBar;

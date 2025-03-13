import React from "react";
import { useNavigate } from "react-router-dom";
import profilePic from "../assets/profile-pic.jpg";
import "../styles/top_bar.css"; // Make sure this file includes the above CSS

const TopBar = () => {
  const navigate = useNavigate();

  const handleEditProfile = () => {
    navigate("/edit-profile");
  };

  const handleLogout = () => {
    localStorage.clear();
    navigate("/login");
  };

  return (
    <div className="top-bar">
      <div className="project-name">Pipeline Management System</div>

      <div className="user-profile">
        <img src={profilePic} alt="Profile" className="profile-icon" />

        <div className="profile-dropdown">
          <button onClick={handleEditProfile} className="dropdown-item">Edit Profile</button>
          <button onClick={handleLogout} className="dropdown-item">Log Out</button>
        </div>
      </div>
    </div>
  );
};

export default TopBar;

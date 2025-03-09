import React, { useState, useEffect } from "react";
// import "../styles/edit_profile.css";

const EditProfile = () => {
  const [user, setUser] = useState({
    name: "",
    role: "",
  });

  // Simulating fetching user data from backend or local storage
  useEffect(() => {
    const fetchUserData = async () => {
      // Replace this with API call
      const userData = {
        name: "John Doe",
        role: "Admin",
      };
      setUser(userData);
    };

    fetchUserData();
  }, []);

  const handleChange = (e) => {
    setUser({ ...user, [e.target.name]: e.target.value });
  };

  const handleSave = () => {
    alert("Profile updated successfully!");
    // Here, you would send data to the backend
  };

  return (
    <div className="edit-profile-container">
      <h2>Edit Profile</h2>
      <div className="form-group">
        <label>Name:</label>
        <input
          type="text"
          name="name"
          value={user.name}
          onChange={handleChange}
        />
      </div>
      <div className="form-group">
        <label>Role:</label>
        <input
          type="text"
          name="role"
          value={user.role}
          onChange={handleChange}
          disabled
        />
      </div>
      <button className="save-button" onClick={handleSave}>Save Changes</button>
    </div>
  );
};

export default EditProfile;

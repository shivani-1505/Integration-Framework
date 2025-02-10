import React, { useState, useEffect } from 'react';
import { createUser, getUser, updateUser } from '../services/userService';

const UserComponent = () => {
    const [user, setUser] = useState({ name: '', email: '' });
    const [userId, setUserId] = useState(null);
    const [isEditing, setIsEditing] = useState(false);

    useEffect(() => {
        if (userId) {
            fetchUser(userId);
        }
    }, [userId]);

    const fetchUser = async (id) => {
        const fetchedUser = await getUser(id);
        setUser(fetchedUser);
    };

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setUser({ ...user, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (isEditing) {
            await updateUser(userId, user);
        } else {
            await createUser(user);
        }
        resetForm();
    };

    const resetForm = () => {
        setUser({ name: '', email: '' });
        setUserId(null);
        setIsEditing(false);
    };

    return (
        <div>
            <h2>{isEditing ? 'Edit User' : 'Create User'}</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    name="name"
                    value={user.name}
                    onChange={handleInputChange}
                    placeholder="Name"
                    required
                />
                <input
                    type="email"
                    name="email"
                    value={user.email}
                    onChange={handleInputChange}
                    placeholder="Email"
                    required
                />
                <button type="submit">{isEditing ? 'Update' : 'Create'}</button>
            </form>
            <button onClick={resetForm}>Cancel</button>
        </div>
    );
};

export default UserComponent;
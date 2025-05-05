// Verify token and load user data
window.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    
    // If no token, redirect to login page
    if (!token) {
        window.location.href = '/';
        return;
    }
    
    // Verify token with backend
    fetch('/verify-token', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Invalid token');
        }
        return response.json();
    })
    .then(data => {
        // Display user information
        document.getElementById('user-email').textContent = data.email;
        document.getElementById('user-role').textContent = data.role_id;
    })
    .catch(error => {
        console.error('Token verification error:', error);
        // Clear token and redirect to login
        localStorage.removeItem('token');
        window.location.href = '/';
    });
    
    // Set up logout functionality
    document.getElementById('logout-btn').addEventListener('click', () => {
        localStorage.removeItem('token');
        window.location.href = '/';
    });
});
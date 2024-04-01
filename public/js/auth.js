function readURL() {
    const path = (window.location.pathname).includes('login') ? '/auth/login' : '/auth/register';
    const urlParams = new URLSearchParams(window.location.search);
    let error = DOMPurify.sanitize(urlParams.get('error'));
    let success = DOMPurify.sanitize(urlParams.get('success'));

    if (error === "invalid-credentials") {
        error = "Invalid credentials";
    } else if (error !== "") {
        error = "Unknown error";
    }

    if (error) {
        showToast({
            icon: 'error',
            text: error
        })
    }

    if (success === "logged-out") {
        success = "Logged out successfully";
    }

    if (success) {
        showToast({
            icon: 'success',
            text: success,
        })
    }

    window.history.replaceState({}, document.title, path);
}

readURL();
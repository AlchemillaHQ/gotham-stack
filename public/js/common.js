function showToast(options) {
    const defaultOptions = {
        toast: true,
        position: 'bottom',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true
    };

    Swal.fire({
        ...defaultOptions,
        ...options,
    });
}

document.body.addEventListener('htmx:responseError', function (event) {
    const reason = event.detail.xhr.response;
    showToast({
        icon: 'error',
        text: reason,
    });
});

document.body.addEventListener("showMessage", function(evt){
    showToast({
        icon: evt.detail.level,
        text: evt.detail.message,
    });
})

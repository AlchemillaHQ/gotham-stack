async function showToast(options) {
    switch (options.type) {
        case 'error':
            window.notyf.open({
                type: 'error',
                message: options.text,
                duration: options.duration || 2500,
            });
            break;
        case 'success':
            window.notyf.open({
                type: 'success',
                message: options.text,
                duration: options.duration || 2500,
            });
            break;
        case 'info':
            window.notyf.open({
                type: 'info',
                message: options.text,
                duration: options.duration || 2500,
            })
            break;
        default:
            window.notyf.open({
                type: 'info',
                message: options.text,
                duration: options.duration || 2500,
            })
            break;
    }
}

document.body.addEventListener('htmx:responseError', function (event) {
    const reason = event.detail.xhr.response;
    showToast({
        type: 'error',
        text: reason,
    });
});

document.body.addEventListener("showMessage", function(evt){
    showToast({
        type: evt.detail.level,
        text: evt.detail.message,
    });
})
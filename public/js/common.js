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
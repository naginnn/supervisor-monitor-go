window.auth = () => {
    return () => {

        const usr = $('#usr').val();
        const pwd = $('#pwd').val();
        const data64 = btoa(`${usr}:${pwd}`);
        const headers = {"Cache-Control": "no-cache, no-store, must-revalidate", // HTTP 1.1.
            "Pragma": "no-cache", "Expires": 0, "Authorization": "Basic " + data64}
        $.ajax({
            url: '/token',
            method: 'get',
            dataType: 'json',
            headers: headers,
            success: function (data) {
                if (data.redirect) {
                    window.location.href = data.redirect;
                } else {

                }
            },
            error (e) {
                console.log(e);
            }
        });
    }
}

(function ($) {
    $('#sign-in').click(auth());
})($);
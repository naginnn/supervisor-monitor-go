window.clearLogs = () => {
    return () => {
        let arr = window.location.href.split('/');
        let url = location.origin + "/logs/clear/" + arr.at(-1)
        $.ajax({
            url: url,
            method: 'post',
            dataType: 'json',
            success: function (data) {
                window.location.reload();
            }
        });
    }
}

window.updateLogs = () => {
    return () => {
        window.location.reload();
    }
}

(function ($) {
    $('#clear-log').click(clearLogs());
    $('#update-log').click(updateLogs());
})($);
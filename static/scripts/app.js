window.controlProcesses = (command) => {

    return () => {
        $('#loading').css("display", "");
        let arr = window.location.href.split('/');
        let url = location.origin + "/control/processes/process/" + arr.at(-1) + "/" + command

    $.ajax({
        url: url,
        method: 'post',
        dataType: 'json',
        success: function (data) {
            $('#loading').css("display", "none");

            window.location.reload();
        }
    });
}}

(function ($) {
    $('#start-app-btn').click(controlProcesses('start'));
    $('#stop-app-btn').click(controlProcesses('stop'));
    $('#restart-app-btn').click(controlProcesses('restart'));
    $('#kill-pid').click(controlProcesses('kill'));
})($);
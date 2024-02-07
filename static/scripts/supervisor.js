window.controlSupervisor = (command) => {
    const token = localStorage.getItem('token');
    return () => {
        let data = {}
        let config = ""
        if (command === "save_config") {
            config = $('#text-config').text();

            console.log(config);
        }
        data = JSON.stringify({ "config": config})

        $('#loading').css("display", "");
        $.ajax({
            url: '/control/config/' + command,         /* Куда пойдет запрос */
            method: 'post',             /* Метод передачи (post или get) */
            dataType: 'json',          /* Тип данных в ответе (xml, json, script, html). */
            data: data,
            headers: {"Content-Type": "Application/json"},
            success: function (data) {   /* функция которая будет выполнена после успешного запроса.  */
                $('#loading').css("display", "none");
                location.reload();
            },
            error: ()=> {
                $('#loading').css("display", "none");
            }
        });
    }
}



(function ($) {
    $('#shutdown').click(controlSupervisor('shutdown'));
    $('#restart').click(controlSupervisor('restart'));
    $('#clear-log').click(controlSupervisor('clear_log'));
    $('#reload-config').click(controlSupervisor('reload_config'));
    $('#shutdown-and-apply-config').click(controlSupervisor('shutdown_apply_config'));
    $('#kill-all-python-processes').click(controlSupervisor('kill_all_python_processes'));
    $('#save-config').click(controlSupervisor('save_config'));
})($);
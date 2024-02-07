window.controlProcesses = (command) => {
    return () => {
        const headers = {"Cache-Control": "no-cache, no-store, must-revalidate", // HTTP 1.1.
            "Pragma": "no-cache", "Expires": 0}
        $('#loading').css("display", "");
        $.ajax({
            url: '/control/processes/' + command,         /* Куда пойдет запрос */
            method: 'post',             /* Метод передачи (post или get) */
            dataType: 'json',          /* Тип данных в ответе (xml, json, script, html). */
            headers: headers, // HTTP 1.0.
            success: function (data) {   /* функция которая будет выполнена после успешного запроса.  */
                $('#loading').css("display", "none");

            },
            302: function (response) {
                window.location.href = "/auth";
            },
            error: (response)=> {
                alert("Команды не доступны");
                window.location.reload();
                $('#loading').css("display", "none");
            },

        });
    }
}

window.control = () => {
    return () => {
        window.location.replace("/config" );
    }

}
const updateAppsState = () => {
    $.ajax({
        url: '/processes',         /* Куда пойдет запрос */
        method: 'get',             /* Метод передачи (post или get) */
        dataType: 'json',          /* Тип данных в ответе (xml, json, script, html). */
        success: function (data) {   /* функция которая будет выполнена после успешного запроса.  */
            $('#spr_state').text("supervisord:" + data.state);
            let apps = data.apps
            $('#apps-body').empty();
            for (key in apps) {
                console.log(apps[key]);
                var row = $('<tr>');
                var hr = $('<a>').attr("href", document.URL + "/" + apps[key]["Name"]);
                hr.attr("class", "nav-link").text(apps[key]["Name"]);
                row.append($('<td>').html(hr));
                row.append($('<td>').html(apps[key]["StateName"]));
                row.append($('<td>').html(apps[key]["Pid"]));
                $('#apps-body').append(row);
            }
        }, error() {
        }
    });
}

(function ($) {

    $('#start-all-apps-btn').click(controlProcesses('start'));
    $('#stop-all-apps-btn').click(controlProcesses('stop'));
    $('#restart-all-apps-btn').click(controlProcesses('restart'));
    $('#supervisor-settings').click(control());
    // $('#supervisor-settings').click(controlProcesses('restart_all'));

    updateAppsState();
    setInterval(updateAppsState, 2000);
})($);


//<p style="color: #c2880b;" className="my-0"><input style="color: #c2880b;" className="my-1" value="{{key}}"/></p>-->
//<p style="color: #000000;" class="my-0"><input style="color: #000000;" class="my-1" value="{{value}}"/></p>-->
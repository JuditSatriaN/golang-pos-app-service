$(function() {
    let baseURL = $('#baseURL').text();
    $.ajax({
        'method': "GET",
        'url': baseURL + "svc/users",
        'contentType': 'application/json',
        'beforeSend': function() {
            $('#wrapper').hide()
            $('#loader').show()
        },
        'complete': function() {
            $("#loader").hide();
            $('#wrapper').show()
        },
    }).done( function(data) {
        jQuery.each(data, function(index, item) {
            // action html template
            let actionHtml = '<div style="text-align: center;">' +
                '<a class="add"><i class="material-icons">&#xE03B;</i></a>' +
                '<a class="edit"><i class="material-icons">&#xE254;</i></a>' +
                '<a class="delete"><i class="material-icons">&#xE872;</i></a></div>'

            // isAdmin html template
            let isAdmin = '<input class="form-check-input" type="checkbox" onclick="return false">'
            if (item["is_admin"]){
                isAdmin = '<input class="form-check-input" type="checkbox" checked onclick="return false">'
            }

            $("#dataTable").DataTable().row.add([
                item["user_id"],
                item["user_name"],
                item["full_name"],
                item["password"],
                '<div class="form-check" style="text-align: center;">' + isAdmin + '</div>',
                actionHtml,
            ]).draw();
        });
    })
})

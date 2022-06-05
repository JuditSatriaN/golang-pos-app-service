$(function() {
    let baseURL = $('#baseURL').text();
    $.ajax({
        'method': "GET",
        'url': baseURL + "svc/members",
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

            $("#dataTable").DataTable().row.add([
                item["id"],
                item["name"],
                item["phone"],
                actionHtml,
            ]).draw();
        });
    })
})

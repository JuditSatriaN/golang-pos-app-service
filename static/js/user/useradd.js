$(function() {
    // action html template
    let actionHtml = '<div style="text-align: center;">' +
        '<a class="add"><i class="material-icons">&#xE03B;</i></a>' +
        '<a class="edit"><i class="material-icons">&#xE254;</i></a>' +
        '<a class="delete"><i class="material-icons">&#xE872;</i></a></div>'

    // Append table with add row form on add new button click
    $(".add-new").on( "click", function(){
        $(this).attr("disabled", "disabled");
        const row = '<tr>' +
            '<td><input type="text" class="form-control" name="user_id" id="user_id"></td>' +
            '<td><input type="text" class="form-control" name="user_name" id="user_name"></td>' +
            '<td><input type="text" class="form-control" name="full_name" id="full_name"></td>' +
            '<td><input type="text" class="form-control" name="password" id="password"></td>' +
            '<td><div class="form-check" style="text-align: center;"><input class="form-check-input" type="checkbox" value=""></div></td>' +
            '<td style="text-align: center;">' + actionHtml + '</td>' +
            '</tr>';
        $("table").prepend(row);
        $("table tbody tr").eq(0).find(".add, .edit").toggle();
    });
});
$(function() {
    // Edit row on edit button click
    $(document).on("click", ".edit", function(e){
        e.preventDefault();
        let rowObj = $(this).parents("tr")

        let col1Obj = rowObj.find("td:nth-child(1)")
        col1Obj.html('<input id="user_id" name="user_id" type="text" class="form-control" value="' + col1Obj.text() + '">')

        let col2Obj = rowObj.find("td:nth-child(2)")
        col2Obj.html('<input id="user_name" name="user_name" type="text" class="form-control" value="' + col2Obj.text() + '">')

        let col3Obj = rowObj.find("td:nth-child(3)")
        col3Obj.html('<input id="full_name" name="full_name" type="text" class="form-control" value="' + col3Obj.text() + '">')

        let col4Obj = rowObj.find("td:nth-child(4)")
        col4Obj.html('<input id="password" name="password" type="password" class="form-control" value="' + col4Obj.text() + '">')

        let col5Obj = rowObj.find("td:nth-child(5)")
        const checkBoxObj = col5Obj.children().children();
        if (checkBoxObj.is(':checked')){
            col5Obj.html('<div class="form-check" style="text-align: center;">' +
                '<input id="is_admin" name="is_admin" type="checkbox" checked="checked" class="form-check-input" value="' + checkBoxObj.val() + '">' +
                '</div>')
        }else{
            col5Obj.html('<div class="form-check" style="text-align: center;">' +
                '<input id="is_admin" name="is_admin" type="checkbox" class="form-check-input" value="' + checkBoxObj.val() + '">' +
                '</div>')
        }

        $(this).parents("tr").find(".add, .edit").toggle();
        $(".add-new").attr("disabled", "disabled");
    });
});
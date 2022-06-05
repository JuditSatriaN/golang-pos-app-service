$(function() {
    // Edit row on edit button click
    $(document).on("click", ".edit", function(){
        let parentObj = $(this).parents("tr")

        let row1Obj = parentObj.find("td:nth-child(1)")
        row1Obj.html('<input id="user_id" name="user_id" type="text" class="form-control" value="' + row1Obj.text() + '">')

        let row2Obj = parentObj.find("td:nth-child(2)")
        row2Obj.html('<input id="user_name" name="user_name" type="text" class="form-control" value="' + row2Obj.text() + '">')

        let row3Obj = parentObj.find("td:nth-child(3)")
        row3Obj.html('<input id="full_name" name="full_name" type="text" class="form-control" value="' + row3Obj.text() + '">')

        let row4Obj = parentObj.find("td:nth-child(4)")
        row4Obj.html('<input id="password" name="password" type="password" class="form-control" value="' + row4Obj.text() + '">')

        let row5Obj = parentObj.find("td:nth-child(5)")
        const checkBoxObj = row5Obj.children().children();
        if (checkBoxObj.is(':checked')){
            row5Obj.html('<div class="form-check" style="text-align: center;">' +
                '<input id="is_admin" name="is_admin" type="checkbox" checked="checked" class="form-check-input" value="' + checkBoxObj.val() + '">' +
                '</div>')
        }else{
            row5Obj.html('<div class="form-check" style="text-align: center;">' +
                '<input id="is_admin" name="is_admin" type="checkbox" class="form-check-input" value="' + checkBoxObj.val() + '">' +
                '</div>')
        }

        $(this).parents("tr").find(".add, .edit").toggle();
        $(".add-new").attr("disabled", "disabled");
    });
});
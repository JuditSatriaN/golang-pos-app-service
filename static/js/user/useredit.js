$(function() {
    // Edit row on edit button click
    $(document).on("click", ".edit", function(){
        let parentObj = $(this).parents("tr")

        let row1Obj = parentObj.find("td:nth-child(1)")
        row1Obj.html('<input type="text" class="form-control" value="' + row1Obj.text() + '">')

        let row2Obj = parentObj.find("td:nth-child(2)")
        row2Obj.html('<input type="text" class="form-control" value="' + row2Obj.text() + '">')

        let row3Obj = parentObj.find("td:nth-child(3)")
        row3Obj.html('<input type="text" class="form-control" value="' + row3Obj.text() + '">')

        let row4Obj = parentObj.find("td:nth-child(4)")
        row4Obj.html('<input type="text" class="form-control" value="' + row4Obj.text() + '">')

        let row5Obj = parentObj.find("td:nth-child(5)")
        const checkBoxObj = row5Obj.children().children();
        if (checkBoxObj.is(':checked')){
            row5Obj.html('<div class="form-check" style="text-align: center;">' +
                '<input type="checkbox" checked="checked" class="form-check-input" value="' + checkBoxObj.val() + '">' +
                '</div>')
        }else{
            row5Obj.html('<div class="form-check" style="text-align: center;">' +
                '<input type="checkbox" class="form-check-input" value="' + checkBoxObj.val() + '">' +
                '</div>')
        }

        $(this).parents("tr").find(".add, .edit").toggle();
        $(".add-new").attr("disabled", "disabled");
    });
});
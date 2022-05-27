$(function() {
    // Add row on add button click
    $(document).on("click", ".add", function(){
        const input = $(this).parents("tr").find('input[type="text"]');
        input.each(function(){
            $(this).parent("td").html($(this).val());
        });
        $(this).parents("tr").find(".add, .edit").toggle();

        let row5Obj = $(this).parents("tr").find("td:nth-child(5)")
        const checkBoxObj = row5Obj.children().children();
        row5Obj.html('<div class="form-check" style="text-align: center;">' +
            '<input type="checkbox" checked="checked" class="form-check-input" value="' + checkBoxObj.val() + '"' +
            'onclick="return false"></div>');
    });
});
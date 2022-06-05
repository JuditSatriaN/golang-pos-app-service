$(function() {
    // Edit row on edit button click
    $(document).on("click", ".edit", function(e){
        e.preventDefault();
        let rowObj = $(this).parents("tr")

        let col1Obj = rowObj.find("td:nth-child(1)")
        col1Obj.html('<input id="id" name="id" type="text" class="form-control" value="' + col1Obj.text() + '">')

        let col2Obj = rowObj.find("td:nth-child(2)")
        col2Obj.html('<input id="name" name="name" type="text" class="form-control" value="' + col2Obj.text() + '">')

        let col3Obj = rowObj.find("td:nth-child(3)")
        col3Obj.html('<input id="phone" name="phone" type="text" class="form-control" value="' + col3Obj.text() + '">')

        $(this).parents("tr").find(".add, .edit").toggle();
        $(".add-new").attr("disabled", "disabled");
    });
});
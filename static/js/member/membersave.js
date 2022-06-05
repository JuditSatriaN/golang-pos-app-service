$(function() {
    // Add row on add button click
    $(document).on("click", ".add", function(e){
        e.preventDefault();
        let jsonObject = {}
        let baseURL = $('#baseURL').text()
        let rowObj = $(this).parents("tr")

        let col1Obj = rowObj.find("td:nth-child(1)")
        jsonObject["id"] = col1Obj.children().val().trim()
        if (jsonObject["id"] === ""){
            alert("ID tidak boleh kosong");
            return
        }

        let col2Obj = rowObj.find("td:nth-child(2)")
        jsonObject["name"] = col2Obj.children().val().trim()

        let col3Obj = rowObj.find("td:nth-child(3)")
        jsonObject["phone"] = col3Obj.children().val().trim()

        $.ajax({
            type: "POST",
            async: false,
            contentType: 'application/json',
            url: baseURL + "svc/member/upsert",
            data: JSON.stringify(jsonObject),
        }).then(function (res) {
            col1Obj.html(res["id"]);
            col2Obj.html(res["name"]);
            col3Obj.html(res["phone"]);
            rowObj.find(".add, .edit").toggle();
            alertify.success("Data berhasil disimpan");
        }).catch(function (a) {
            alertify.error("Error : " + a.responseText);
        });

    });
});
$(function() {
    // Add row on add button click
    $(document).on("click", ".add", function(e){
        e.preventDefault();
        let jsonObject = {}
        let baseURL = $('#baseURL').text()
        let rowObj = $(this).parents("tr")

        let col1Obj = rowObj.find("td:nth-child(1)")
        jsonObject["user_id"] = col1Obj.children().val()
        if (jsonObject["user_id"] === ""){
            alert("User ID tidak boleh kosong");
            return
        }

        let col2Obj = rowObj.find("td:nth-child(2)")
        jsonObject["user_name"] = col2Obj.children().val()

        let col3Obj = rowObj.find("td:nth-child(3)")
        jsonObject["full_name"] = col3Obj.children().val()

        let col4Obj = rowObj.find("td:nth-child(4)")
        jsonObject["password"] = col4Obj.children().val()
        if (jsonObject["password"] === ""){
            alert("Password tidak boleh kosong");
            return
        }

        let col5Obj = rowObj.find("td:nth-child(5)")
        let checkBoxObj = col5Obj.children().children();
        jsonObject["is_admin"] = checkBoxObj.is(':checked')

        $.ajax({
            type: "POST",
            contentType: 'application/json',
            url: baseURL + "svc/user/upsert",
            data: JSON.stringify(jsonObject),
        }).then(function (res) {
            alert("Data berhasil disimpan");
            col1Obj.html(res["user_id"]);
            col2Obj.html(res["user_name"]);
            col3Obj.html(res["full_name"]);
            col4Obj.html(res["password"]);

            // isAdmin html template
            let isAdmin = '<input class="form-check-input" type="checkbox" onclick="return false">'
            if (res["is_admin"]){
                isAdmin = '<input class="form-check-input" type="checkbox" checked onclick="return false">'
            }
            col5Obj.html('<div class="form-check" style="text-align: center;">' + isAdmin + '</div>');
            rowObj.find(".add, .edit").toggle();
        }).catch(function (a) {
            alert("Error : " + a.responseText);
        });

    });
});
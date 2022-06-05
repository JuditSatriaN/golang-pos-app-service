$(function() {
    // Delete row on delete button click
    $(document).on("click", ".delete", function(){
        if (!confirm('Apakah anda yakin ingin menghapus data ini?')) {
            return
        }
        let jsonObject = {}
        let baseURL = $('#baseURL').text();
        let rowObj = $(this).parents("tr")

        let col1Obj = rowObj.find("td:nth-child(1)")
        jsonObject["user_id"] = col1Obj.text()
        $.ajax({
            type: "POST",
            contentType: 'application/json',
            url: baseURL + "svc/user/delete",
            data: JSON.stringify(jsonObject),
        }).then(function (res) {
            rowObj.remove();
            alert("Data berhasil dihapus")
            $(".add-new").removeAttr("disabled");
        }).catch(function (a) {
            alert("Error : " + a.responseText);
        });
    });
});
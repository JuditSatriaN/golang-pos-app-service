$(function() {
    // Delete row on delete button click
    $(document).on("click", ".delete", function(e){
        e.preventDefault()
        let baseURL = $('#baseURL').text();
        let rowObj = $(this).parents("tr")

        if (rowObj.find("td:nth-child(1)").text() === "") {
            rowObj.remove();
            $(".add-new").removeAttr("disabled");
            return
        }

        alertify.confirm('Dialog Konfirmasi', 'Apakah anda yakin ingin menghapus data ini?', function(){ deleteData(baseURL, rowObj) }, null);

    });
});

function deleteData(baseURL, rowObj){
    let jsonObject = {}

    let col1Obj = rowObj.find("td:nth-child(1)")
    jsonObject["id"] = col1Obj.text()
    $.ajax({
        type: "POST",
        async: false,
        contentType: 'application/json',
        url: baseURL + "svc/member/delete",
        data: JSON.stringify(jsonObject),
    }).then(function (res) {
        rowObj.remove();
        alertify.success("Data berhasil dihapus")
        $(".add-new").removeAttr("disabled");
    }).catch(function (a) {
        alertify.error("Error : " + a.responseText);
    });
}
function NumSystemConvertion() {
    let source_num = document.getElementById("source_num").value
    let source_pow = document.getElementById("source_pow").value
    let result_pow = document.getElementById("result_pow").value
    document.getElementById("result_num").innerHTML = parseInt(source_num, source_pow).toString(result_pow)
}

function CreateMultTable() {
    let div_table = document.getElementById("mult_table")
    let mult_table = "<table>"
    let rows = document.getElementById("table_rows").value
    let cols = document.getElementById("table_cols").value
    let color = document.getElementById("table_color").value
    if (rows > 20) {
        alert('Число строк не может быть больше 20!')
        return
    }
    if (cols > 20) {
        alert('Число столбцов не может быть больше 20!')
        return
    }
    for (i = 0; i <= rows; i++) {
        mult_table += '<tr>'    
        for (j = 0; j <= cols; j++) {
            var elem = '<td>' + i * j + '</td>'
            if (i == 0) {
                elem = `<td style="background-color: ${color};">` + j + '</td>'
            }
            if (j == 0) {
                elem = `<td style="background-color: ${color};">` + i + '</td>'
            }
            mult_table += elem
        }
        mult_table += '</tr>'
    }
    mult_table += "</table>"
    div_table.innerHTML = mult_table
}


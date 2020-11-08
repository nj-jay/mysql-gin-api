$(function() {

				getDate()
				function getDate() {
					var temp = '';
					
					$.ajax({
						type: "get",
						url: "http://nj-jay.com:8080/api/v1/books",
						dataType: "json",
          	 			//contentType: "application/json",
						//async: false,
						success: function(res) {
							//alert(res)
							var list = res
							for(var $i = 0; $i < list.length; $i++) {
								temp +=
									'<tr>' +
									'<td>' + list[$i].id + '</td>' +
									'<td>' + list[$i].name + '</td>' +
									'<td>' + list[$i].price+ '</td>' +
									'</tr>';
							}
							$("#jsonTable tr:not(:first)").html(""); /*  除了第一行tr的内容，其余内容清空，防止获取重复的数据  */
							$("#jsonTable").append(temp);
						}
					});
				}
			});
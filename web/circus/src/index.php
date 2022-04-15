 <html>
<head>
<title>Equal or not equal?</title>
<link rel="stylesheet" href="/bootstrap.min.css"  crossorigin="anonymous">

</head>
<body>
<div class="container">
<div class="jumbotron" style="top: 50px; position: absolute; text-align: center; width: 70%; ">
<br></br>
<h3> А ты знаешь C достаточно хорошо? </h3>
<br></br>
<form action="index.php" method="post">
	<input type="text" name="pass" class="form-control"> <br></br>
	<input type="submit" class="btn btn-danger"><br></br>
</form>
<?php 
	
	$pass = $_POST['pass'];

	if(strcmp('2218043148asdfasdf123asdf1234rasdf345ewasdfasd!@#asdf123df245f4', $pass)) {
		echo "<h3> Похоже что нет =( </h3>";
	}
	else {
		echo " <h3> Молодец!!  YetiCTF{g00d_g4m3} </h3>";
	}

?>

</div>
</div>
</body>
</html> 

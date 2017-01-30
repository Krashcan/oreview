<html>
<head>
    <meta charset="UTF-8">
    <title>Movies</title>
    <link rel="stylesheet" type="text/css" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <link rel="stylesheet" type="text/css" href="static/styles.css">
    <script src="semantic.min.js"></script>
</head>
<body>
        {{range $i,$f:= .}}
       
        <div class="container">
            <div class="row">
            <div class="col-md-4">
                <img src="{{$f.Image}}" class="img-thumbnail">
            </div>
            <div class="col-md-8">
                <h3>{{$f.Title}}</h3><h3>{{$f.Year}}</h3><h2>{{$f.Rating}}</h2>
                <h6>{{$f.Runtime}}</h6><h6> {{$f.Genre}}</h6>
                <br>
                <p>{{$f.Description}}</p>
                <p>{{$f.Awards}}</p>
            </div>
            
        </div>
        </div>
   
    {{end}}
    
</body>
</html>

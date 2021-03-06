<html>
<head>
    <meta charset="UTF-8">
    <title>Movies</title>
    <link type="text/css" rel="stylesheet" href="/static/semantic/semantic.min.css">
    <link rel="stylesheet" type="text/css" href="static/styles.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
    <script src="/static/semantic/semantic.min.js"></script>
    <script src="semantic.min.js"></script>
</head>
<body>
    <header>
            <div class="logo">
                <p>Oreview</p>
            </div>
    </header>
    <div class="ui four doubling stackable special cards">
        {{range $i,$f:= .}}
    
      <div class="ui fluid card">
        <div class="blurring dimmable image">
          <div class="ui dimmer">
            <div class="content">
              <div class="center">
                <h3>{{$f.Description}}</h3>
              </div>
            </div>
          </div>
          <img src="{{$f.Image}}">
        </div>
        <div class="content">
          <p class="header">{{$f.Title}}</p>
          <p class="header"><b>IMDb {{$f.Rating}}</b></p>
          <div class="meta">
            <span class="date">{{$f.Year}}</span>
          </div>
        </div>
        <div class="extra content">
            <p>
                <i class="comment"></i>
                {{$f.Genre}}
            </p>
            <p>
                <i class="hourglass end icon"></i>
                {{$f.Runtime}}
            </p>
            
        </div>
      </div>
        {{end}}
    </div>
       
    <script>
    $('.special.cards .image').dimmer({
        on: 'hover'
    });
    </script>
    <footer>
        <i class="star icon"></i> and Contribute on <a href="https://github.com/Krashcan/oreview/">Github</a><br>
        <i class="copyright icon"></i>Licensed by <a href="https://github.com/Krashcan/oreview/blob/master/License.md">MIT</a>
    </footer>
</body>
</html>

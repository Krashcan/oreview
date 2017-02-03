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
    <div class="ui special cards">
        {{range $i,$f:= .}}
      <div class="card">
        <div class="blurring dimmable image">
          <div class="ui dimmer">
            <div class="content">
              <div class="center">
                <p >IMDb rating</p>
                <h2>{{$f.Rating}}</h2>
              </div>
            </div>
          </div>
          <img src="{{$f.Image}}">
        </div>
        <div class="content">
          <a class="header">{{$f.Title}}</a>
          <div class="meta">
            <span class="date">{{$f.Year}}</span>
          </div>
        </div>
        <div class="extra content">
          <a>
            <i class="hourglass end icon"></i>
            {{$f.Runtime}}
          </a>
          <a>
            <i class="comment"></i>
            {{$f.Description}}
          </a>      
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

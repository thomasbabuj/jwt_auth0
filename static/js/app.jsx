
var App = React.createClass({
  componentWillMount : function(){
    this.setState({idToken: null})
  },
  render : function() {
    if( this.state.idToken ) {
      return (<LoggedIn idToken={this.state.idToken} />);
    } else {
      console.log( 'dsdsfd');
      return (<Home />)
    }
  }
});

var Home = React.createClass({
  render : function() {
    return (
      <div className="container">
        <div className="cols-xs-12 jumbotron text-center">
          <h1>We R VR</h1>
          <p>Provide valuable feedback to VR experience developers.</p>
          <a className="btn btn-primary btn-lg btn-block">Sign In</a>
        </div>
      </div>
    );
  }
});

var LoggedIn = React.createClass({
  getInitialState : function() {
    return {
      profile : null,
      products: null
    }
  },
  render : function() {
    if (this.state.profile) {
      return (
        <div className="col-lg-12">
          <span className="pull-right">
            {this.state.profile.nickname}
            <a onClick={this.logout}>Logout</a>
          </span>
          <h2>Welcome to We R VR</h2>
          <p>Below you'll find the latest games that need feedback. Please provide
          honest feedback so developers can make the best games.
          </p>
          <div className="row">
            {this.state.products.map(function(product,i){
              return <Product key="{i}" product={product} />
            })}
          </div>
        </div>
      );
    } else {
      return (<div>Loading...</div>);
    }
  }
});

var Product = React.createClass({
  upvote : function() {
  },
  downvote : function() {
  },
  getInitialState : function() {
      return {
        voted: null
      }
  },
  render : function() {
    return(
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">
            {this.props.product.name}
            <span className="pull-right">{this.state.voted}</span>
          </div>
          <div className="panel-body">
            {this.props.product.Description}
          </div>
          <div className="panel-footer">
            <a onClick="{this.upvote}" className="btn btn-default">
              <span className="glyphicon glyphicon-thumbs-up"></span>
            </a>
            <a onClick="{this.downvote}" className="btn btn-default">
              <span className="glyphicon glyphicon-thumbs-down"></span>
            </a>
          </div>
        </div>
      </div>
    )
  }
});

ReactDOM.render(<App />, document.getElementById('app'))

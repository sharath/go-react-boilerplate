import React from 'react'

class Demo extends React.Component {
    constructor() {
        super();
        this.state = {users: []};
    }

    componentDidMount() {
        fetch('api/v1/users').then((response) => response.json()).then((js) => this.setState({users: js}));
    }

    render() {
        let users = this.state.users;
        return (
            <ul>
                {users.map(user =>
                    <li>{user}</li>
                )}
            </ul>
        );
    }

}

export default Demo;
import React from 'react';
import PropTypes from 'prop-types';
import { Button } from './Button';
import { Input } from './Input';
import { Link } from './Link';

function validName(name){
	return name.length > 2;
}

function validEmail(email){
	return /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email)
}

function validPassword(password){
	return password.length > 7;
}

export const Register = ({ onSubmit, onClickLogin }) => {

	const [user, setUser] = React.useState({username: "", email: "", password: "", password2: ""})
	const [valid, setValid] = React.useState({username: true, email: true, password: true, password2: true})
	const err = "border-2 border-red-500"

	const validatePassword = (pass1, pass2)  => {
		let pass1Valid = validPassword(pass1)
		setValid({...valid, 
			password: pass1Valid,
			password2: !pass1Valid || pass2 == pass1
		});
	}


	return (
		<div className={`bg-secondary p-6 rounded-md sm:w-400 `}>
			<span className="text-2xl font-bold" >Register</span>
			<Input
				className={`mt-3 ${valid.username ? "" : err}`}
				placeholder="Username"
				value={user.username}
				onChange={(e) => {
					setUser({...user, username: e.target.value}); 
					setValid({...valid, username: validName(e.target.value)});
				}}
			/>
			<Input
				className={`mt-3 ${valid.email ? "" : err}`}
				placeholder="Email"
				value={user.email}
				onChange={(e) => {
					setUser({...user, email: e.target.value});  
					setValid({...valid, email: validEmail(e.target.value)}); 
				}}
			/>
			<Input
				className={`mt-3 ${valid.password ? "" : err}`}
				placeholder="Password"
				type={"password"}
				value={user.password}
				onChange={(e) => {
					setUser({...user, password: e.target.value}); 
					validatePassword(e.target.value, user.password2);
				}}
			/>
			<Input
				className={`mt-3 ${valid.password2 ? "" : err}`}
				placeholder="Confirm Password"
				type={"password"}
				value={user.password2}
				onChange={(e) => {
					setUser({...user, password2: e.target.value}); 
					validatePassword(user.password, e.target.value);
				}}
			/>
			<div className="flex justify-between mt-4">
				<Button
					size="medium"
					color="primary"
					disabled={false}
					loading={false}
					onClick={() => {onSubmit(user)}}
				>
					Register
				</Button>
				<span className="my-auto">
					<Link onClick={() => {onClickLogin()}} >Already have an account</Link>
				</span>
			</div>
		</div>
	);
};

Register.propTypes = {
	onSubmit: PropTypes.func,
	onClickLogin: PropTypes.func,
};
import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import { Button } from './Button';
import { Input } from './Input';
import { Link } from './Link';

function validName(name){
	return name.length > 2;
}

function validEmail(email){
	/* eslint-disable-next-line */
	return /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email)
}

function validPassword(password){
	return password.length > 7;
}

const err = "border-2 border-red-500"

export const Register = ({ onSubmit, onClickLogin }) => {

	const [user, setUser] = React.useState({username: "", email: "", password: "", password2: ""})
	const [valid, setValid] = React.useState({username: true, email: true, password: true, password2: true})

	useEffect(() => { validateFields(true) }, [user, setUser])

	const validateFields = (isAllowEmpty) => {

		let validPass = validPassword(user.password)
		let validPass2 = !validPass || user.password === user.password2

		let update = {
			username: validName(user.username) || (!user.username && isAllowEmpty), 
			email: validEmail(user.email) || (!user.email && isAllowEmpty), 
			password: validPass || (!user.password && isAllowEmpty), 
			password2: validPass2, 
		}

		setValid(update)

		return update
	}

	const formSubmit = () => {

		let fields = validateFields(false)

		for (let field in fields) {
			if (!fields[field]) return
		}

		onSubmit(user)
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
				}}
			/>
			<Input
				className={`mt-3 ${valid.email ? "" : err}`}
				placeholder="Email"
				value={user.email}
				onChange={(e) => {
					setUser({...user, email: e.target.value});  
				}}
			/>
			<Input
				className={`mt-3 ${valid.password ? "" : err}`}
				placeholder="Password"
				type={"password"}
				value={user.password}
				onChange={(e) => {
					setUser({...user, password: e.target.value}); 
				}}
			/>
			<Input
				className={`mt-3 ${valid.password2 ? "" : err}`}
				placeholder="Confirm Password"
				type={"password"}
				value={user.password2}
				onChange={(e) => {
					setUser({...user, password2: e.target.value}); 
				}}
			/>
			<div className="flex justify-between mt-4">
				<Button
					size="medium"
					color="primary"
					disabled={false}
					loading={false}
					onClick={() => {formSubmit()}}
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
import React from 'react';
import PropTypes from 'prop-types';
import { Button } from './Button';
import { Input } from './Input';


const Link = ({ children, ...props }) => {
	return (
		<span className="text-sm underline text-secondary-text cursor-pointer select-none" {...props} >{children}</span>
	);
};

export const Register = ({ onSubmit, onClickLogin }) => {

	const [user, setUser] = React.useState({username: "", email: "", password: "", password2: ""})
	const [valid, setValid] = React.useState({username: true, email: true, password: true, password2: true})

	return (
		<div className={`bg-secondary p-6 rounded-md sm:w-400 `}>
			<span className="text-2xl font-bold" >Register</span>
			<Input
				className="mt-3"
				placeholder="Username"
				value={user.username}
				onChange={(e) => {setUser({...user, username: e.target.value})}}
			/>
			<Input
				className="mt-2"
				placeholder="Email"
				value={user.email}
				onChange={(e) => {setUser({...user, email: e.target.value})}}
			/>
			<Input
				className="mt-2"
				placeholder="Password"
				type={"password"}
				value={user.password}
				onChange={(e) => {setUser({...user, password: e.target.value})}}
			/>
			<Input
				className="mt-2"
				placeholder="Confirm Password"
				type={"password"}
				value={user.password2}
				onChange={(e) => {setUser({...user, password2: e.target.value})}}
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
					<Link onClick={() => {onClickLogin()}} >Already have account</Link>
				</span>
			</div>
		</div>
	);
};

Register.propTypes = {
	onSubmit: PropTypes.func,
	onClickLogin: PropTypes.func,
};
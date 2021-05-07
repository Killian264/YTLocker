import React from 'react';
import PropTypes from 'prop-types';
import { Button } from './Button';
import { Input } from './Input';


const Link = ({ children, ...props }) => {
	return (
		<span className="text-sm underline text-secondary-text cursor-pointer select-none" {...props} >{children}</span>
	);
};

export const Login = ({ onSubmit, onClickRegister }) => {

	const [user, setUser] = React.useState({email: "", password: ""})

	return (
		<div className={`bg-secondary p-6 rounded-md sm:w-400 `}>
			<span className="text-2xl font-bold" >Login</span>
			<Input
				className="mt-3"
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
			<div className="flex justify-between mt-4">
				<Button
					size="medium"
					color="primary"
					disabled={false}
					loading={false}
					onClick={() => {onSubmit(user)}}
				>
					Login
				</Button>
				<span className="my-auto">
					<Link onClick={() => {onClickRegister()}} >Register</Link>
				</span>
			</div>
		</div>
	);
};

Login.propTypes = {
	onSubmit: PropTypes.func,
	onClickRegister: PropTypes.func,
};
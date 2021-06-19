import React from "react";
import { Button } from "./Button";
import { Input } from "./Input";
import { Link } from "./Link";

export interface LoginProps {
	className?: string;
	onSubmit: (email: string, password: string) => void;
	onClickRegister: () => void;
}

export const Login: React.FC<LoginProps> = ({ className, onSubmit, onClickRegister }) => {
	const [user, setUser] = React.useState({
		email: "",
		password: "",
	});

	return (
		<div className={`${className} bg-primary-700 p-10 rounded-md sm:w-96 `}>
			<span className="text-2xl font-bold">Login</span>
			<Input
				className="mt-3"
				placeholder="Email"
				value={user.email}
				onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
					setUser({ ...user, email: e.target.value });
				}}
				onKeyDown={(e: React.KeyboardEvent<HTMLInputElement>) => {
					if (e.key === "Enter") {
						onSubmit(user.email, user.password);
					}
				}}
			/>
			<Input
				className="mt-2"
				placeholder="Password"
				type={"password"}
				value={user.password}
				onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
					setUser({ ...user, password: e.target.value });
				}}
				onKeyDown={(e: React.KeyboardEvent<HTMLInputElement>) => {
					if (e.key === "Enter") {
						onSubmit(user.email, user.password);
					}
				}}
			/>
			<div className="flex justify-between mt-4">
				<Button
					size="medium"
					color="primary"
					disabled={false}
					loading={false}
					onClick={() => {
						onSubmit(user.email, user.password);
					}}
				>
					Login
				</Button>
				<span className="my-auto">
					<Link
						onClick={() => {
							onClickRegister();
						}}
					>
						Create an account instead
					</Link>
				</span>
			</div>
		</div>
	);
};

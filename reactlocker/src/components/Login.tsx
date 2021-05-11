import React from "react";
import { Button } from "./Button";
import { Input } from "./Input";
import { Link } from "./Link";
import { UserLogin } from "../shared/types";

export interface LoginProps {
	className?: string;
	onSubmit: (user: UserLogin) => void;
	onClickRegister: () => void;
}

export const Login: React.FC<LoginProps> = ({
	className,
	onSubmit,
	onClickRegister,
}) => {
	const [user, setUser] = React.useState<UserLogin>({
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
			/>
			<Input
				className="mt-2"
				placeholder="Password"
				type={"password"}
				value={user.password}
				onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
					setUser({ ...user, password: e.target.value });
				}}
			/>
			<div className="flex justify-between mt-4">
				<Button
					size="medium"
					color="primary"
					disabled={false}
					loading={false}
					onClick={() => {
						onSubmit(user);
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

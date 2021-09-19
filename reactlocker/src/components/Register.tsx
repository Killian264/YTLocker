import React, { useEffect } from "react";
import { Button } from "./Button";
import { Input } from "./Input";
import { Link } from "./Link";
import { validateFields } from "../shared/validation";
import { UserRegister } from "../hooks/api/useUserRegister";

export interface RegisterProps {
	className?: string;
	onSubmit: (user: UserRegister) => void;
	onClickLogin: () => void;
}

const err = "border-2 border-red-500";

export const Register: React.FC<RegisterProps> = ({ onSubmit, onClickLogin }) => {
	const [user, setUser] = React.useState<UserRegister>({
		username: "",
		email: "",
		password: "",
		password2: "",
	});
	const [valid, setValid] = React.useState({
		username: true,
		email: true,
		password: true,
		password2: true,
	});

	useEffect(() => {
		let update = validateFields(user, true);
		setValid(update);
	}, [user, setUser]);

	const formSubmit = () => {
		let fields = validateFields(user, false);
		setValid(fields);

		// TODO learn how to iterate instead
		if (!fields.username || !fields.email || !fields.password || !fields.password2) {
			return;
		}

		onSubmit(user);
	};

	return (
		<div className={`bg-primary-700 p-10 rounded-md sm:w-96 `}>
			<span className="text-2xl font-bold">Register</span>
			<Input
				className={`mt-3 ${valid.username ? "" : err}`}
				placeholder="Username"
				value={user.username}
				onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
					setUser({ ...user, username: e.target.value });
				}}
			/>
			<Input
				className={`mt-3 ${valid.email ? "" : err}`}
				placeholder="Email"
				value={user.email}
				onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
					setUser({ ...user, email: e.target.value });
				}}
			/>
			<Input
				className={`mt-3 ${valid.password ? "" : err}`}
				placeholder="Password"
				type={"password"}
				value={user.password}
				onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
					setUser({ ...user, password: e.target.value });
				}}
			/>
			<Input
				className={`mt-3 ${valid.password2 ? "" : err}`}
				placeholder="Confirm Password"
				type={"password"}
				value={user.password2}
				onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
					setUser({ ...user, password2: e.target.value });
				}}
			/>
			<div className="flex justify-between mt-4">
				<Button
					size="medium"
					color="primary"
					disabled={false}
					loading={false}
					onClick={() => {
						formSubmit();
					}}
				>
					Register
				</Button>
				<span className="my-auto">
					<Link
						onClick={() => {
							onClickLogin();
						}}
					>
						Already have an account
					</Link>
				</span>
			</div>
		</div>
	);
};

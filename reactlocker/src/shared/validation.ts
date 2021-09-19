import { UserRegister } from "../hooks/api/useUserRegister";

export const validName = (name: string): boolean => {
	return name.length > 2;
};

export const validEmail = (email: string): boolean => {
	/* eslint-disable-next-line */
	return /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email);
};

export const validPassword = (password: string): boolean => {
	return password.length > 7;
};

export const validateFields = (user: UserRegister, isAllowEmpty: boolean) => {
	let validPass = validPassword(user.password);
	let validPass2 = !validPass || user.password === user.password2;

	let update = {
		username: validName(user.username) || (!user.username && isAllowEmpty),
		email: validEmail(user.email) || (!user.email && isAllowEmpty),
		password: validPass || (!user.password && isAllowEmpty),
		password2: validPass2,
	};

	return update;
};

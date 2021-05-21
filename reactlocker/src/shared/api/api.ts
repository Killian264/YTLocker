import { UserLogin, UserRegister } from "../types";
import { DROPLET_BASE } from "../env";

export interface ApiResponse {
	success: boolean;
	message: string;
}

export interface ApiResponse2 {
	status: number;
	message: string;
}

export interface InternalResponse extends ApiResponse {
	data: any;
}

interface UserLoginResponse extends ApiResponse {
	bearer: string;
}

const Login = (user: UserLogin) => {
	const url = DROPLET_BASE + "/user/login";

	const options = {
		method: "POST",
		headers: {},
		body: JSON.stringify(user),
	};

	return fetch(url, options)
		.then((res) => res.json())
		.then((data) => ParseApiResponse(data))
		.then((data) => ParseLoginResponse(data));
};

const Register = (user: UserRegister) => {
	const url = DROPLET_BASE + "/user/register";

	const options = {
		method: "POST",
		headers: {},
		body: JSON.stringify(user),
	};

	return fetch(url, options)
		.then((res) => res.json())
		.then((data) => ParseApiResponse(data))
		.then((data) => ParseRegisterResponse(data));
};

export const API = {
	UserLogin: Login,
	UserRegister: Register,
};

const ParseLoginResponse = (data: InternalResponse): UserLoginResponse => {
	if (!data.success) {
		return {
			...data,
			bearer: "",
		};
	}

	return {
		...data,
		bearer: data.success ? data.data.Bearer : "",
	};
};

const ParseRegisterResponse = (data: InternalResponse): ApiResponse => {
	return {
		...data,
	};
};

const ParseApiResponse = (data: any): InternalResponse => {
	return {
		success: data.Status === 200,
		message: data.Message,
		data: data.Data,
	};
};

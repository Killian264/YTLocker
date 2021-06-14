import axios from "axios";
import { useLocalStorage } from "./useLocalStorage";

export const useBearer = (initial: string): [string, (bearer: string) => void] => {
	const [cookie, setCookie] = useLocalStorage("bearer", newBearer(initial));

	const setBearer = (bearer: string) => {
		setCookie(newBearer(bearer));
	};

	if (isExpired(new Date(cookie.timestamp))) {
		return [initial, setBearer];
	}

	return [cookie.value, setBearer];
};

const newBearer = (value: string) => {
	return { value: value, timestamp: Date.now() };
};

const isExpired = (date: Date) => {
	const twentyHours = 20 * 60 * 60 * 1000;
	const expires = new Date().valueOf() - twentyHours;
	return date.valueOf() < expires;
};

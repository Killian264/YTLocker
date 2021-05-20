import { useHistory } from "react-router";
import { useLocalStorage } from "./useLocalStorage";

export const useBearer = (initial: string): [string, (bearer: string) => void] => {
	let history = useHistory();
	const [cookie, setCookie] = useLocalStorage("bearer", {
		bearer: initial,
		date: Date.now(),
	});

	const setBearer = (bearer: string) => {
		let cookie = {
			bearer: bearer,
			date: Date.now(),
		};

		setCookie(cookie);
	};

	if (isExpired(new Date(cookie.date))) {
		history.push("/login");
	}

	return [cookie.bearer, setBearer];
};

const isExpired = (date: Date) => {
	const twentyHours = 20 * 1000 * 60 * 60;
	const expires = Date.now() - twentyHours;
	return date < new Date(expires);
};

import { useQuery } from "react-query";
import { useHistory } from "react-router";
import { DROPLET_BASE } from "../env";
import { useBearer } from "../hooks/useBearer";
import { ApiResponse2 } from "./api";

export interface LatestFetchResponse extends ApiResponse2 {
	videos: number[];
}

export const useLatestVideos = (): [boolean, number[]] => {
	const [bearer] = useBearer("");
	let history = useHistory();

	const { isLoading, isError, data } = useQuery(["latest"], () => fetchLatest(bearer));

	if (isLoading || isError || data === undefined) {
		return [true, []];
	}

	if (data.status === 401) {
		history.push("/login");
	}

	if (data.videos === null) {
		return [true, []];
	}

	return [false, data.videos];
};

const fetchLatest = async (bearer: string): Promise<LatestFetchResponse> => {
	const url = DROPLET_BASE + `/playlist/videos/latest`;

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	const res = await (await fetch(url, options)).json();

	let latest: LatestFetchResponse = {
		status: res.Status,
		message: res.Message,
		videos: [],
	};

	if (latest.status === 200) {
		latest.videos = res.Data;
	}

	return latest;
};

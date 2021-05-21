import { useQuery } from "react-query";
import { useHistory } from "react-router";
import { DROPLET_BASE } from "../env";
import { useBearer } from "../hooks/useBearer";
import { Channel } from "../types";
import { ApiResponse2 } from "./api";

export interface ChannelFetchResponse extends ApiResponse2 {
	channel: Channel | null;
}

export const useChannel = (id: number): [boolean, Channel | null] => {
	const [bearer] = useBearer("");
	let history = useHistory();

	const { isLoading, isError, data } = useQuery(["channel", id], () => fetchChannel(bearer, id));

	if (isLoading || isError || data === undefined) {
		return [true, null];
	}

	if (data.status === 401) {
		history.push("/login");
	}

	if (data.channel === null) {
		return [true, null];
	}

	return [false, data.channel];
};

const fetchChannel = async (bearer: string, id: number): Promise<ChannelFetchResponse> => {
	const url = DROPLET_BASE + `/channel/${id}`;

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	const res = await (await fetch(url, options)).json();

	let channel: ChannelFetchResponse = {
		status: res.Status,
		message: res.Message,
		channel: null,
	};

	if (channel.status === 200) {
		channel.channel = parseChannel(res.Data);
	}

	return channel;
};

const parseChannel = (data: any): Channel => {
	let thumbnail = data.Thumbnails[data.Thumbnails.length - 1];

	let parsed: Channel = {
		id: data.ID,
		youtubeId: data.YoutubeID,
		thumbnailUrl: thumbnail.URL,
		title: data.Title,
		description: data.Description,
		created: new Date(Date.parse(data.CreatedAt)),
		videos: data.Videos,
	};

	return parsed;
};

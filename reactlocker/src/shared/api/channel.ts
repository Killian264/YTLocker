import { DROPLET_BASE } from "../env";
import { Channel } from "../types";
import { ApiResponse } from "./api";

export interface ChannelFetchResponse extends ApiResponse {
	channel: Channel | null;
}

export const fetchChannel = async (bearer: string, id: number): Promise<ChannelFetchResponse> => {
	const url = DROPLET_BASE + `/channel/${id}`;

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	const res = await (await fetch(url, options)).json();

	const data = res.Data;

	let channel: ChannelFetchResponse = {
		success: res.Status == 200,
		message: res.Message,
		channel: null,
	};

	if (channel.success) {
		channel.channel = parseChannel(data);
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

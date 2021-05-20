import { DROPLET_BASE } from "../env";
import { Video } from "../types";
import { ApiResponse } from "./api";

export interface VideoFetchResponse extends ApiResponse {
	video: Video | null;
}

export const fetchVideo = async (bearer: string, id: number): Promise<VideoFetchResponse> => {
	const url = DROPLET_BASE + `/video/${id}`;

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	const res = await (await fetch(url, options)).json();

	const data = res.Data;

	let video: VideoFetchResponse = {
		success: res.Status == 200,
		message: res.Message,
		video: null,
	};

	if (video.success) {
		video.video = parseVideo(data);
	}

	return video;
};

const parseVideo = (data: any): Video => {
	let thumbnail = data.Thumbnails[data.Thumbnails.length - 1];

	let parsed: Video = {
		id: data.ID,
		youtubeId: data.YoutubeID,
		thumbnailUrl: thumbnail.URL,
		title: data.Title,
		description: data.Description,
		created: new Date(Date.parse(data.CreatedAt)),
	};

	return parsed;
};

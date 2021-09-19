import React from "react";
import { usePlaylistList } from "../hooks/api/usePlaylistList";
import { VideoListItemController } from "./VideoListItemController";
import { Card } from "../components/Card";
import { useVideoListLatest } from "../hooks/api/useVideoListLatest";
import { LoadingList } from "../components/LoadingList";

export interface VideoListLatestControllerProps {
	className?: string;
	limit: number;
}

export const VideoListLatestController: React.FC<VideoListLatestControllerProps> = ({ className, limit }) => {
	const [isLoadingPlaylists, playlists] = usePlaylistList();
	const [isLoadingVideos, vidoes] = useVideoListLatest();

	if (isLoadingPlaylists || isLoadingVideos || playlists == null) {
		return <LoadingList limit={10}></LoadingList>;
	}

	let count = 0;
	let list: JSX.Element[] = [];

	vidoes.forEach((videoId) => {
		playlists.forEach((playlist) => {
			if (!playlist.videos.includes(videoId)) {
				return;
			}

			if (count >= limit) {
				return;
			}
			count++;

			list.push(
				<VideoListItemController
					key={count}
					videoId={videoId}
					playlist={playlist}
				></VideoListItemController>
			);
		});
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<span className="leading-none text-2xl font-semibold">Videos</span>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

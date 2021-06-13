import React from "react";
import { Card } from "../components/Card";
import { Playlist } from "../shared/types";
import { VideoListItemController } from "./VideoListItemController";

export interface PlaylistVideoListControllerProps {
	className?: string;
	limit: number;
	playlist: Playlist;
}

export const PlaylistVideoListController: React.FC<PlaylistVideoListControllerProps> = ({
	className,
	limit,
	playlist,
}) => {
	const list = playlist.videos.map((id, index) => {
		if (index >= limit) {
			return "";
		}
		return (
			<VideoListItemController key={index} videoId={id} playlist={playlist}></VideoListItemController>
		);
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Videos</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

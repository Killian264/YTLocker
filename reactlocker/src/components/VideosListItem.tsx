import React from "react";
import { Video } from "../shared/types";
import { Link } from "./Link";
import { Arrow } from "./Svg";

export interface VideoListItemProps {
	className?: string;
	video: Video;
}

export const VideoListItem: React.FC<VideoListItemProps> = ({
	video,
	className,
}) => {
	return (
		<div
			className={`${className} hover:bg-primary-600 rounded-md flex justify-between`}
		>
			<div className="flex p-1">
				<img
					src={video.thumbnail}
					alt="Logo"
					width="140"
					height="120"
					className="rounded-lg"
				/>
				<div className="pl-3 flex flex-col">
					<span className="text-lg font-semibold">{video.title}</span>
					<Link
						className="text-accent text-lg ml-0.5"
						href={video.url}
						target="_blank"
					>
						Youtube Link
					</Link>
				</div>
			</div>
			<div className="mr-3 my-auto select-none">
				<Arrow size={24}></Arrow>
			</div>
		</div>
	);
};

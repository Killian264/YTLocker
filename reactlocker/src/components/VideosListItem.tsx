import React from "react";
import { Color, Video } from "../shared/types";
import { ColorBadge } from "./ColorBadge";
import { RightArrow } from "./Svg";

export interface VideoListItemProps {
	className?: string;
	video: Video;
	url: string;
	color: Color;
}

export const VideoListItem: React.FC<VideoListItemProps> = ({ video, className, url, color }) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between cursor-pointer`;
	const imageSize = "md:h-20 sm:h-16 h-16";
	const textSize = "sm:text-md text-md";

	const open = () => {
		window.open(url, "_blank");
	};

	let trimmedTitle = video.title.length > 30 ? video.title.substring(0, 33) + "..." : video.title;

	return (
		<div className={css} onClick={open}>
			<div className="flex p-1">
				<img
					src={video.thumbnailUrl}
					alt="Thumbnail"
					className={`${imageSize} rounded-lg object-cover`}
				/>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold`}>{trimmedTitle}</span>
					<span className="text-sm text-primary-text-200">{video.created.toDateString()}</span>
					<div>
						<ColorBadge className="mt-2" color={color}></ColorBadge>
					</div>
				</div>
			</div>
			<div className="mr-2 my-auto select-none">
				<RightArrow size={24}></RightArrow>
			</div>
		</div>
	);
};

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
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between cursor-pointer`;

	const imageSize = "md:h-20 sm:h-16 h-16";

	const textSize = "sm:text-lg text-md";

	const open = () => {
		window.open(video.url, "_blank");
	};

	return (
		<div className={css} onClick={open}>
			<div className="flex p-1">
				<img
					src={video.thumbnail}
					alt="Logo"
					className={`rounded-lg object-cover ${imageSize}`}
				/>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold`}>
						{video.title}
					</span>
					<Link
						className={`${textSize} text-accent ml-0.5`}
						href={video.url}
						target="_blank"
					>
						Youtube Link
					</Link>
				</div>
			</div>
			<div className="mr-2 my-auto select-none">
				<Arrow size={24}></Arrow>
			</div>
		</div>
	);
};

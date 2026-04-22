import 'package:dio/dio.dart';
import 'package:retrofit/retrofit.dart';
import 'package:json_annotation/json_annotation.dart';

part 'api_client.g.dart';

@RestApi()
abstract class EdgeNodeClient {
  factory EdgeNodeClient(Dio dio, {String baseUrl}) = _EdgeNodeClient;

  @GET("/api/v1/feed")
  Future<List<Article>> getFeed();
}

@JsonSerializable()
class Article {
  @JsonKey(name: 'id')
  final String id;
  @JsonKey(name: 'rephrased_headline_dv')
  final String rephrasedHeadlineDv;
  @JsonKey(name: 'category')
  final String category;
  @JsonKey(name: 'visual_source')
  final String visualSource;
  @JsonKey(name: 'visual_url')
  final String visualURL;
  @JsonKey(name: 'subscriber_briefing')
  final String subscriberBriefing;
  @JsonKey(name: 'subscriber_explainer')
  final String subscriberExplainer;
  @JsonKey(name: 'is_breaking')
  final bool isBreaking;
  @JsonKey(name: 'is_premium')
  final bool isPremium;
  @JsonKey(name: 'is_ad')
  final bool isAd;

  Article({
    required this.id,
    required this.rephrasedHeadlineDv,
    required this.category,
    required this.visualSource,
    required this.visualURL,
    required this.subscriberBriefing,
    required this.subscriberExplainer,
    required this.isBreaking,
    required this.isPremium,
    required this.isAd,
  });

  factory Article.fromJson(Map<String, dynamic> json) => _$ArticleFromJson(json);
  Map<String, dynamic> toJson() => _$ArticleToJson(this);
}
